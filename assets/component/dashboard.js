import { html } from "https://unpkg.com/lit-html?module";

const dashboard = () => {
  return html`
    <toast-noti
      titulo="Titulo"
      message="Prueba mensaje"
      duration="50000"
    />
  `;
};
export { dashboard };
