document.addEventListener('DOMContentLoaded', (event) => {
    // function openModal() {
    //     const errorTarget = document.getElementById("htmx-alert");
    //     if (errorTarget) {
    //         errorTarget.setAttribute("hidden", "true");
    //         errorTarget.innerText = "";
    //     }

    //     const modalContent = document.getElementById("modal-content");
    //     if (modalContent) {
    //         modalContent.style.display = "flex";
    //     } else {
    //         console.log("missing el with id: modal-content")
    //     };
    // }

    // function closeModal() {
    //     const modalContent = document.getElementById("modal-content");
    //     if (modalContent) {
    //         modalContent.style.display = "none";
    //     } else {
    //         console.log("missing el with id: modal-content")
    //     };
    // }

    // window.addEventListener('keydown', function (event) {
    //     if (event.key === 'Escape') {
    //         closeModal();
    //     }
    // });

    // window.openModal = openModal;
    // window.closeModal = closeModal;

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
    )
});
