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

    const handleAjax = (element, form) => {
        const method = element.hasAttribute('hx-get') ? 'GET' :
                       element.hasAttribute('hx-post') ? 'POST' : 'DELETE';
        const endpoint = element.getAttribute('hx-get') || 
                         element.getAttribute('hx-post') || 
                         element.getAttribute('hx-delete');
        const targetSelectors = element.getAttribute('hx-swap');
        const targetElements = targetSelectors ? targetSelectors.split(',').map(selector => document.querySelector(selector.trim())) : [];

        let fetchOptions = { method };
        if (method === 'POST' || method === 'DELETE') {
            const formData = new FormData(form);
            fetchOptions.body = formData;
        }

        fetch(endpoint, fetchOptions)
            .then(response => response.text())
            .then(data => {
                let extractedHTML = parseString(data, targetElements);

                targetElements.forEach((targetElement) => {
                    if (targetElement) {
                        const responseElement = extractedHTML.get(targetElement.id);

                        if (responseElement) {
                            targetElement.innerHTML = responseElement;
                        }
                    }
                });
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

    const addSearchItemListener = (searchItem) => {
        searchItem.addEventListener('dblclick', event => {
            const inputs = searchItem.getElementsByTagName('input');
            
            for (let input of inputs) {
                input.value = "";
            }
        });
    };

    const searchItems = document.getElementById('items-search');
    addSearchItemListener(searchItems);

    const observer = new MutationObserver((mutationsList) => {
        for (let mutation of mutationsList) {
            if (mutation.type === 'childList') {
                mutation.addedNodes.forEach(node => {
                    if (node.nodeType === Node.ELEMENT_NODE) {
                        const elements = node.querySelectorAll('button[hx-get], button[hx-post], button[hx-delete], a[hx-get], a[hx-post], a[hx-delete]');
                        addEventListeners(elements);

                        if (node.id === 'items-search') {
                            addSearchItemListener(node);
                        } else {
                            const searchItem = node.querySelector('#items-search');
                            if (searchItem) {
                                addSearchItemListener(searchItem);
                            }
                        }
                    }
                });
            }
        }
    });

    observer.observe(document.body, { childList: true, subtree: true });
});
