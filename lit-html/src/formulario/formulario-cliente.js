import {LitElement, html} from 'lit-element';
import '@vaadin/vaadin-text-field/vaadin-text-field.js';
import '@vaadin/vaadin-form-layout/vaadin-form-layout.js';
import '@vaadin/vaadin-button/vaadin-button.js';

export class FormularioCliente extends LitElement {
  static get properties() {
    return {
      _id: {
        type: Number,
      },
      cliente: {
        type: String,
      },
      numeroFactura: {
        type: String,
      },
      ordenDeCompra: {
        type: String,
      },
    };
  }

  constructor() {
    super();
    //Nombre del cliente
    this.cliente = '';
    //Numero de factura
    this.numeroFactura = '';
    //Numero de orden de compra
    this.ordenDeCompra = '';
  }

  fijarCliente({target: {value}}) {
    this.cliente = value;
  }

  fijarNumeroFactura({target: {value}}) {
    this.numeroFactura = value;
  }

  fijarOrdenDeCompra({target: {value}}) {
    this.ordenDeCompra = value;
  }
  //TODO Probar el generador de id's
  guardarFormulario() {
    const orden = {
      cliente: this.cliente,
      numeroFactura: this.numeroFactura,
      ordenDeCompra: this.ordenDeCompra,
    };
    //Si este formulario cuenta con un id, significa que esta es una actualizacion
    const event = new CustomEvent(this._id ? 'actualizar' : 'guardar', {
      detail: {
        orden,
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
          @change=${this.fijarCliente}
        ></vaadin-text-field>

        <vaadin-text-field
          label="Factura"
          @change=${this.fijarNumeroFactura}
        ></vaadin-text-field>

        <vaadin-text-field
          label="Orden de compra"
          @change=${this.fijarOrdenDeCompra}
        ></vaadin-text-field>

        <vaadin-button @click=${this.guardarFormulario}>Guardar</vaadin-button>
      </vaadin-form-layout>
    `;
  }
}

customElements.define('formulario-cliente', FormularioCliente);
