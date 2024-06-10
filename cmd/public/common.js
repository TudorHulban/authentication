document.addEventListener('DOMContentLoaded', (event) => {
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
});
