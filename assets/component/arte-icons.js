export class ArteIcons extends HTMLSpanElement {
  constructor() {
    super();
    this.link = document.createElement("link");
    this.link.href = "https://fonts.googleapis.com/icon?family=Material+Icons";
    this.link.rel = "stylesheet";
    this.link.type = "text/css";
  }
  connectedCallback() {
    this.appendChild(this.link);
    this.classList.add("material-icons");
    this.appendChild(document.createTextNode(this.icon));
  }
  get icon() {
    return this.getAttribute("icon");
  }
  set icon(newValue) {
    this.setAttribute("icon", newValue);
  }
}

customElements.define("arte-icons", ArteIcons, { extends: "span" });
