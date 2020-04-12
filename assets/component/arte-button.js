const style = `
	<style>
		::host{
			background:RED			
		}
	</style>
`;

export class ArteButton extends HTMLButtonElement {
  constructor() {
    super();
  }
}

customElements.define("arte-button", ArteButton, { extends: "button" });
