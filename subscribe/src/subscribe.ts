
export class Subscribe extends HTMLElement {

    connectedCallback() {
        this.collectState();
        const observer = new MutationObserver((mutations) => {
            mutations.forEach((mutation) => {
                if (mutation.addedNodes.length > 0) {
                    this.collectState();
                }
            })
        })

        observer.observe(this, { childList: true })
    }

    private collectState() {
        const form = this.querySelector("form") as HTMLFormElement;
        if (!form) {
            return;
        }

        const button = form.querySelector("button[type=submit]") as HTMLButtonElement;
        const input = form.querySelector("input[type=email]") as HTMLButtonElement;

        if (!button) {
            throw new Error("form without button");
        }

        button.toggleAttribute("disabled", false);
        input.focus();

        form.addEventListener("htmx:beforeRequest", function() {
            button.toggleAttribute("disabled", true);
        })

        form.addEventListener("htmx:beforeSwap", function() {
            console.log("BEFORE SWAP");
            button.toggleAttribute("false", true);
        })
    }
}

