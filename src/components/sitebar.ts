class MySidebar extends HTMLElement {
    async connectedCallback() {
        const resp = await fetch("/src/components/sidebar.html");
        this.innerHTML = await resp.text();
        console.log("123");
    }
}

customElements.define("my-sidebar", MySidebar);
