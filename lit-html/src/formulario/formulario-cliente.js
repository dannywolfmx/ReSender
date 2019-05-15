import {LitElement, html} from 'lit-element';
import '@vaadin/vaadin-text-field/vaadin-text-field.js';
import '@vaadin/vaadin-form-layout/vaadin-form-layout.js';
import '@vaadin/vaadin-button/vaadin-button.js';

export class FormularioCliente extends LitElement {
  static get properties() {
    return {
      orden: {
        type: Object,
      },
    };
  }

  constructor() {
    super();
    this.orden = {};
  }

  fijarCliente({target: {value}}) {
    this.orden.cliente = value;
  }

  fijarNumeroFactura({target: {value}}) {
    this.orden.numeroFactura = value;
  }

  fijarOrdenDeCompra({target: {value}}) {
    this.orden.ordenDeCompra = value;
  }

  guardarFormulario() {
    //Si este formulario cuenta con un id, significa que esta es una actualizacion

    const event = new CustomEvent(this.orden._id ? 'actualizar' : 'guardar', {
      detail: {
        orden: this.orden,
        error: false,
      },
    });

    this.dispatchEvent(event);
  }
  //https://developer.mozilla.org/en-US/docs/Web/API/HTML_Drag_and_Drop_API/File_drag_and_drop
  guardarArchivo(evento) {
    evento.preventDefault();
    console.log(evento);
    if (evento) {
    }
  }

  render() {
    return html`
      <vaadin-form-layout @drop="${this.guardarArchivo}">
        <vaadin-text-field
          label="cliente"
          value="${this.orden.cliente}"
          @change=${this.fijarCliente}
        ></vaadin-text-field>

        <vaadin-text-field
          label="Factura"
          value="${this.orden.numeroFactura}"
          @change=${this.fijarNumeroFactura}
        ></vaadin-text-field>

        <vaadin-text-field
          label="Orden de compra"
          value=${this.orden.ordenDeCompra}
          @change=${this.fijarOrdenDeCompra}
        ></vaadin-text-field>

        <vaadin-button @click=${this.guardarFormulario}>Guardar</vaadin-button>
      </vaadin-form-layout>
    `;
  }
}

customElements.define('formulario-cliente', FormularioCliente);
