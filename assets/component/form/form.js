import { html, render } from "https://unpkg.com/lit-html?module";

import "../bootstrap-input.js";
import "../arte-button.js";
class MyForm extends HTMLElement {
  constructor() {
    super();
    this.inputs = [];
    this.inputHide = [];
  }
  setNameForm(nameForm) {
    this.nameForm = nameForm;
  }

  setInput(inputs) {
    inputs.map((i) => {
      if (i.hide) {
        this.inputHide.push(i);
      } else {
        this.inputs.push(i);
      }
    });
  }
  connectedCallback() {
    this._update();
  }

  //Implements by the Class child
  submitCallback(data, form) {}

  _submit(e) {
    e.preventDefault();
    let form = document.forms[this.nameform];
    console.log(form);
    console.log(this.nameform);
    let data = new FormData(form);
    this.submitCallback(data, form);
  }

  _input(label, type, name) {
    return html`
      <bootstrap-input
        label=${label}
        type=${type}
        name=${name}
        slot="search"
      ></bootstrap-input>
    `;
  }
  _template() {
    return html`
      <form name=${this.nameform}>
        <div class="form-group">
          ${this.inputs.map((i) => this._input(i.label, i.type, i.name))}
        </div>
        <button
          is="arte-button"
          type="submit"
          class="btn btn-primary"
          @click=${this._submit}
        >
          Crear
        </button>
      </form>
    `;
  }

  _update() {
    render(this._template(), this, { eventContext: this });
  }
}

customElements.define("my-form", MyForm);

export { MyForm };
