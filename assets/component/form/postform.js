import { MyForm } from "./form.js";

export class POSTForm extends MyForm {
  constructor() {
    super();
  }

  submitCallback(data, form) {
    let datos = Object.fromEntries(data);
    this.inputs.map((i) => {
      if (i.type === "number") {
        datos[i.name] = Number(datos[i.name]);
      }
    });
    this.inputHide.map((i) => {
      datos[i.name] = i.value;
    });

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

  attributeChangedCallback(attr, oldval, newval) {
    if (attr === "inputs") {
      this.setInputs(JSON.parse(newval));
    }
  }

  setInputs(inputs) {
    super.setInput(inputs);
  }
  get nameform() {
    return this.getAttribute("nameform");
  }

  set nameform(newValue) {
    this.setAttribute("nameform", newValue);
  }
  get url() {
    return this.getAttribute("url");
  }

  set url(newValue) {
    this.setAttribute("url", newValue);
  }

  static get observedAttributes() {
    return ["inputs"];
  }

  _update() {
    super._update();
  }
}

customElements.define("post-form", POSTForm);
