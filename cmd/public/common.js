document.
    addEventListener('DOMContentLoaded', (event) => {
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
                    const parser = new DOMParser();
                    const doc = parser.parseFromString(data, 'text/html');

                    console.log(doc);
                    
                    targetElements.forEach((targetElement, index) => {
                        if (targetElement) {
                            const responseElement = doc.querySelector(`#${targetElement.id}`);
                            if (responseElement) {
                                targetElement.innerHTML = responseElement.innerHTML;
                            }
                        }
                    });
                })
                .catch(error => console.error('Error:', error));
        };

        document.querySelectorAll('button[hx-get], button[hx-post], button[hx-delete], a[hx-get], a[hx-post], a[hx-delete]').forEach(element => {
            element.addEventListener('click', function (event) {
                event.preventDefault();
                const form = element.closest('form') || document.createElement('form'); // Create a dummy form if not inside a form
                handleAjax(element, form);
            });
        });

        document.querySelectorAll('.ajax-form').forEach(form => {
            form.addEventListener('submit', function(event) {
                event.preventDefault();
                const submitButton = form.querySelector('button[type="submit"]');
                handleAjax(submitButton, form);
            });

            form.querySelectorAll('.deleteButton').forEach(deleteButton => {
                deleteButton.addEventListener('click', function(event) {
                    event.preventDefault();
                    handleAjax(deleteButton, form);
                });
            });
        });

        const searchItems = document.getElementById('items-search');
        searchItems.addEventListener(
            'dblclick', event => {
                const inputID = document.getElementById('items-search-id');
                inputID.value = "";

                const inputStatus = document.getElementById('items-search-status');
                inputStatus.value = "";

                const inputName = document.getElementById('items-search-name');
                inputName.value = "";
            }
        );
    });
