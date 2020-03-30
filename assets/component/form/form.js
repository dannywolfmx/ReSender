import { html, render } from "https://unpkg.com/lit-html?module";

class MyForm extends HTMLElement {
  constructor(nameForm, inputs, sumbitEvent) {
    super();
    this.nameForm = nameForm;
    this.inputs = inputs;
    //Function to send data
    this.sumbitEvent = sumbitEvent;
  }
  connectedCallback() {
    this._update();
  }

  _submit(e) {
    e.preventDefault();
    let form = document.forms[this.nameForm];
    let data = new FormData(form);
    this.sumbitEvent(data, form);
  }

  _input(label, type, name) {
    return html`
      <label for="exampleInputEmail1">${label}</label>
      <input type=${type} class="form-control" name=${name} />
    `;
  }
  _template() {
    return html`
      <form name=${this.nameForm}>
        <div class="form-group">
          ${this.inputs.map((i) => this._input(i.label, i.type, i.name))}
        </div>
        <button type="submit" class="btn btn-primary" @click=${this._submit}>
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
