# authentication with htmx

## Resources

```html
https://medium.com/@hhartleyjs/making-a-login-page-with-htmx-go-5acbcc504426
```

### Multiswap

```html
https://github.com/bigskysoftware/htmx-extensions/blob/main/src/multi-swap/README.md
```

### Error Handling

```html
https://xvello.net/blog/htmx-error-handling/
```  

### Code for modal support

Not sure if to use modals though.

```js
    function openModal() {
        const errorTarget = document.getElementById("htmx-alert");
        if (errorTarget) {
            errorTarget.setAttribute("hidden", "true");
            errorTarget.innerText = "";
        }

        const modalContent = document.getElementById("modal-content");
        if (modalContent) {
            modalContent.style.display = "flex";
        } else {
            console.log("missing el with id: modal-content")
        };
    }

    function closeModal() {
        const modalContent = document.getElementById("modal-content");
        if (modalContent) {
            modalContent.style.display = "none";
        } else {
            console.log("missing el with id: modal-content")
        };
    }

    window.addEventListener('keydown', function (event) {
        if (event.key === 'Escape') {
            closeModal();
        }
    });

    window.openModal = openModal;
    window.closeModal = closeModal;
```  

## Code for form support

```js
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
```
