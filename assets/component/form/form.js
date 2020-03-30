import { html, render } from "https://unpkg.com/lit-html?module";

class MyForm extends HTMLElement {
  constructor(nameForm, inputs, url) {
    super();
    this.nameForm = nameForm;
    this.inputHide = [];
    this.inputs = [];
    inputs.map((i) => {
      if (i.hide) {
        this.inputHide.push(i);
      } else {
        this.inputs.push(i);
      }
    });
    //Function to send data
    this.url = url;
    this.sumbitEvent = this._defaultSubmit;
  }
  connectedCallback() {
    this._update();
  }

  _defaultSubmit(data, form) {
    let datos = Object.fromEntries(data);
    this.inputs.map((i) => {
      if (i.type === "number") {
        datos[i.name] = Number(datos[i.name]);
      }
    });
    this.inputHide.map((i) => {
      datos[i.name] = i.value;
    });
    console.log(datos);
    fetch(this.url, {
      method: "POST",
      body: JSON.stringify(datos),
    }).then((r) => {
      if (r.ok) {
        form.reset();
      } else {
        console.log("Error");
      }
    });
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
