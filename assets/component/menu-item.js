import { html, render } from "https://unpkg.com/lit-html?module";
const style = html`
  <style>
    :host {
      display: block;
      padding: 8px 0px;
    }

    :host(:hover) {
      background: #6777ef;
      color: white;
    }
    p {
      margin: 0px;
    }
  </style>
`;
export class MenuItem extends HTMLElement {
  constructor() {
    super();
    this.root = this.attachShadow({ mode: "open" });
  }
  connectedCallback() {
    this._update();
  }
  _template() {
    return html`
      ${style}
      <p>${this.title}</p>
    `;
  }
  _update() {
    render(this._template(), this.root, { eventContext: this });
  }
}
customElements.define("menu-item", MenuItem);
