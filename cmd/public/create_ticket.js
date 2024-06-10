document.body.addEventListener('htmx:afterRequest',
    function (evt) {
        const errorTarget = document.getElementById("htmx-alert");

        if (evt.detail.successful) {
            errorTarget.setAttribute("hidden", "true");
            errorTarget.innerText = "";

            window.alert("Succes - Item created!")

            window.location.replace("/tickets");
        } else if (evt.detail.failed && evt.detail.xhr) {
            console.warn("server error", evt.detail);

            const xhr = evt.detail.xhr;
            let errorMsg = { error: "Unexpected server response." };

            try {
                const parsedResponse = JSON.parse(xhr.responseText);
                if (parsedResponse && parsedResponse.error) {
                    errorMsg = parsedResponse;
                } else {
                    console.error("Parsed response does not contain 'error' property:", parsedResponse);
                }
            } catch (e) {
                console.error("Failed to parse JSON response:", e);
            }

            errorTarget.innerText = `Error from server: ${xhr.status} - ${JSON.stringify(errorMsg)}`;
            errorTarget.removeAttribute("hidden");
        } else {
            console.error("unexpected server error", evt.detail);

            errorTarget.innerText = "Unexpected error, check your connection and try to refresh the page.";
            errorTarget.removeAttribute("hidden");
        }
    }
);