class NavBar extends HTMLElement{
	constructor() {
		super()
		//this.attachShadow({mode:'open'})
	}
	
	connectedCallback(){
		this.template()

	}
	
	template(){
		this.innerHTML = `
			<nav>
			  <a hrefi="/" class="brand">ReSender</a>
			
			  <!-- responsive-->
			  <input id="bmenug" type="checkbox" class="show">
			  <label for="bmenug" class="burger pseudo button">&#8801;</label>
			
			  <div class="menu">
			    <input placeholder="Buscar" />
			  </div>
			</nav>
		`	
	}
}


customElements.define("nav-bar", NavBar)
