import { html, render } from "https://unpkg.com/lit-html?module";

const style = html`
  <style>
    :host {
      display: flex;
      align-items: center;
      justify-content: center;
      margin: 8px;
    }
    input {
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.03);
      padding: 8px;
      border-radius: 6px;
      border: 1px solid #ecececd4;
      margin-left: 16px;
    }
  </style>
`;

export class BootstrapInput extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: "open" });
  }
  connectedCallback() {
    this._update();
  }
  set label(newValue) {
    this.setAttribute("label", newValue);
  }
  get label() {
    return this.getAttribute("label") || "";
  }
  set type(newValue) {
    this.setAttribute("type", newValue);
  }
  get type() {
    return this.getAttribute("type") || "";
  }
  set name(newValue) {
    this.setAttribute("name", newValue);
  }
  get name() {
    return this.getAttribute("name") || "";
  }
  _template() {
    return html`
      ${style}
      <label>${this.label}</label>
      <input type=${this.type} name=${this.name} />
    `;
  }
  _update() {
    render(this._template(), this.root, { eventContext: this });
  }
}

customElements.define("bootstrap-input", BootstrapInput);
