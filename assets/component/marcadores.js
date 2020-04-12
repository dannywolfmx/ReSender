import { html, render } from "https://unpkg.com/lit-html?module";

const style = html`
  <style>
    :host {
      min-height: 224px;
      height: auto;
      display: flex;
      padding: 16px;
    }
    ::slotted(.card) {
      margin-left: 16px;
      width: 448px;
      border: none !important;
      box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3);
      background: white;
    }
  </style>
`;

export class MyMarcadores extends HTMLElement {
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
      <slot name="marcadores"></slot>
    `;
  }
  _update() {
    render(this._template(), this.root, { eventContext: this });
  }
}

customElements.define("my-marcadores", MyMarcadores);
