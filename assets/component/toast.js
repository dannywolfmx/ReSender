import { html, render } from "https://unpkg.com/lit-html?module";

class Toast extends HTMLElement {
  constructor() {
    super();
    this.toast = document.createElement("div");
    this.setAttribute(
      "style",
      "position: relative; min-height:200px; display:block;"
    );
    this.toast.setAttribute(
      "style",
      "position: absolute; top: 20px; right: 20px; min-width:350px;"
    );
    this.toast.setAttribute("class", "toast fade");
    this.appendChild(this.toast);
    this._update();
  }
  connectedCallback() {}

  attributeChangedCallback(attr, oldval, newval) {
    if (attr === "titulo") {
      this.titulo = newval;
      this._update();
    } else if (attr == "message") {
      this.message = newval;
      this._update();
    } else if (attr == "duration") {
      this._duration(newval);
    }
  }

  static get observedAttributes() {
    return ["titulo", "message", "duration"];
  }
  //Determinar cuanto turara esta notificacion
  _duration(miliseconds) {
    this._show();
    console.log("Mostrando");
    setTimeout(this._hide.bind(this), miliseconds);
  }
  _show() {
    //Cambiar el opacity para mostrar el elemento
    this.toast.classList.remove("hide");
    setTimeout(() => {
      this.toast.classList.add("show");
    }, 0.1);
  }

  _hide() {
    this.toast.classList.remove("show");
    setTimeout(() => {
      this.toast.classList.add("hide");
    }, 150);
  }

  _toastHeader() {
    return html`
      <div class="toast-header">
        <img src="..." class="rounded mr-2" alt="..." />
        <strong class="mr-auto">${this.titulo}</strong>
        <small>11 mins ago</small>
        <button
          type="button"
          class="ml-2 mb-1 close"
          data-dismiss="toast"
          aria-label="Close"
          @click=${this._hide}
        >
          <span aria-hidden="true">&times;</span>
        </button>
      </div>
    `;
  }

  _toastBody() {
    return html`
      <div class="toast-body">
        ${this.message}
      </div>
    `;
  }

  _template() {
    return html` ${this._toastHeader()} ${this._toastBody()}`;
  }

  _update() {
    render(this._template(), this.toast, { eventContext: this });
  }
}

customElements.define("toast-noti", Toast);
