document.addEventListener('DOMContentLoaded', (event) => {
    function parseString(inputString, elements) {
        let parts = inputString.split('|');
        let matchedElementsMap = new Map();
        
        for (let part of parts) {
            for (let element of elements) {
                if (part.includes(`id="${element.id}"`)) {
                    matchedElementsMap.set(element.id, part);
                    break;
                }
            }
        }
        
        return matchedElementsMap;
    }

    function showPopover(element, message) {
        const popover = document.createElement('div');
        popover.className = 'popover';
        popover.textContent = message;
        document.body.appendChild(popover);

        const rect = element.getBoundingClientRect();
        popover.style.left = `${rect.left + window.scrollX}px`;
        popover.style.top = `${rect.bottom + window.scrollY}px`;

        setTimeout(() => {
            document.body.removeChild(popover);
        }, 3000);
    }

    const validateRequirements = (element) => {
        const requireAttr = element.getAttribute('hx-require');
        if (requireAttr) {
            const requiredIds = requireAttr.split(',');
            for (let id of requiredIds) {
                const requiredElement = document.querySelector(id.trim());
                if (requiredElement && !requiredElement.value) {
                    showPopover(requiredElement, 'This field is required.');
                    return false;
                }
            }
        }
        return true;
    };

    const handleAjax = (element, form) => {
        if (!validateRequirements(element)) {
            return;
        }

        const method = element.hasAttribute('hx-get') ? 'GET' :
                       element.hasAttribute('hx-post') ? 'POST' : 'DELETE';
        const endpoint = element.getAttribute('hx-get') || 
                         element.getAttribute('hx-post') || 
                         element.getAttribute('hx-delete');
        const targetSelectors = element.getAttribute('hx-swap');
        const targetElements = targetSelectors ? targetSelectors.split(',').map(selector => document.querySelector(selector.trim())) : [];
        const redirectUrl = element.getAttribute('hx-redirect');

        let fetchOptions = { method };
        if (method === 'POST' || method === 'DELETE') {
            const formData = new FormData(form);
            fetchOptions.body = formData;
        }

        fetch(endpoint, fetchOptions)
            .then(response => {
                if (response.ok && redirectUrl) {
                    window.location.href = redirectUrl;
                    return null; // No need to process further
                } else {
                    return response.text();
                }
            })
            .then(data => {
                if (data) {
                    let extractedHTML = parseString(data, targetElements);

                    targetElements.forEach((targetElement) => {
                        if (targetElement) {
                            const responseElement = extractedHTML.get(targetElement.id);

                            if (responseElement) {
                                targetElement.innerHTML = responseElement;
                            }
                        }
                    });
                }
            })
            .catch(error => console.error('Error:', error));
    };

    const addEventListeners = (elements) => {
        elements.forEach(element => {
            element.addEventListener('click', function (event) {
                event.preventDefault();
                const form = element.closest('form') || document.createElement('form'); // Create a dummy form if not inside a form
                handleAjax(element, form);
            });
        });
    };

    document.querySelectorAll('button[hx-get], button[hx-post], button[hx-delete], a[hx-get], a[hx-post], a[hx-delete]').forEach(element => {
        element.addEventListener('click', function (event) {
            event.preventDefault();
            const form = element.closest('form') || document.createElement('form'); // Create a dummy form if not inside a form
            handleAjax(element, form);
        });
    });

    const addDblClickInputClearerListener = (searchItem) => {
        searchItem.addEventListener('dblclick', event => {
            const inputs = searchItem.getElementsByTagName('input');
            
            for (let input of inputs) {
                input.value = "";
            }
        });
    };

    const searchItems = document.getElementById('items-search');
    addDblClickInputClearerListener(searchItems);

    const observer = new MutationObserver((mutationsList) => {
        for (let mutation of mutationsList) {
            if (mutation.type === 'childList') {
                mutation.addedNodes.forEach(node => {
                    if (node.nodeType === Node.ELEMENT_NODE) {
                        const elements = node.querySelectorAll('button[hx-get], button[hx-post], button[hx-delete], a[hx-get], a[hx-post], a[hx-delete]');
                        addEventListeners(elements);

                        if (node.id === 'items-search') {
                            addDblClickInputClearerListener(node);
                        } else {
                            const searchItem = node.querySelector('#items-search');
                            if (searchItem) {
                                addDblClickInputClearerListener(searchItem);
                            }
                        }
                    }
                });
            }
        }
    });

    observer.observe(document.body, { childList: true, subtree: true });
});
