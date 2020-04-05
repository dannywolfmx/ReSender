import { html, render } from "https://unpkg.com/lit-html?module";
const style = html`
  <style>
    :host {
      width: 100%;
      display: block;
      padding: 16px 0px;
    }

    :host(:hover) {
      background: #6777ef;
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
      <slot></slot>
    `;
  }
  _update() {
    render(this._template(), this.root, { eventContext: this });
  }
}
customElements.define("menu-item", MenuItem);
