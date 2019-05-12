import {LitElement, html} from 'lit-element';
import './lista-ordenes.js';
import './formulario/formulario-cliente.js';
export class StartLitElement extends LitElement {
  /**
   * Define properties. Properties defined here will be automatically
   * observed.
   */
  static get properties() {
    return {
      ordenes: {
        type: Array,
      },
    };
  }

  /**
   * In the element constructor, assign default property values.
   */
  constructor() {
    // Must call superconstructor first.
    super();

    //Solicitar lista de ordenes
    this.actualizarLista();
  }
  //Actualizar la lista de ordenes
  actualizarLista() {
    const lista = localStorage.getItem('ordenes');

    this.ordenes = JSON.parse(lista) || [];
  }

  /**
   * Define a template for the new element by implementing LitElement's
   * `render` function. `render` must return a lit-html TemplateResult.
   */
  render() {
    return html`
      <h1>RSender</h1>
      <formulario-cliente @guardar="${this.agregarOrden}"></formulario-cliente>
      <lista-ordenes .lista="${this.ordenes}"></lista-ordenes>
    `;
  }

  //Agregar una nueva orden de compra
  agregarOrden({detail: {orden, error}}) {
    console.log(orden);

    if (error) {
      //Notificar error
      return -1;
    }

    let ordenes = localStorage.getItem('ordenes');
    let _id = localStorage.getItem('_idOrden');
    //Verificar que el id sea valido
    if (!_id) {
      _id = 0;
    } else {
      _id++;
    }

    localStorage.setItem('_idOrden', _id);

    orden._id = _id;

    ordenes = ordenes ? JSON.parse(ordenes) : [];
    ordenes.push(orden);

    localStorage.setItem('ordenes', JSON.stringify(ordenes));
    this.actualizarLista();
  }
}

// Register the element with the browser
customElements.define('start-lit-element', StartLitElement);
