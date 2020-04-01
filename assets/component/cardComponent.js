import { html, render } from "https://unpkg.com/lit-html?module";

class Card extends HTMLElement {
  constructor() {
    super();
    console.log("PruebPruebaa");
    this._update();
  }

  attributeChangedCallback(attr, oldval, newval) {}

  static get observedAttributes() {
    return ["cardTitle"];
  }
  _template() {
    return html`
      <div class="card">
        <div class="card-body">
          <h5 class="card-title">${this.title}</h5>
          <slot> </slot>
        </div>
      </div>
    `;
  }

  _update() {
    render(this._template(), this, { eventContext: this });
  }
}

customElements.define("bootstrap-card", Card);
