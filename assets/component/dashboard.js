import { html } from "https://unpkg.com/lit-html?module";

const dashboard = () => {
  return html`<toast-noti
    titulo="Titulo"
    message="Prueba mensaje"
    duration="5000"
  />`;
};
export { dashboard };
