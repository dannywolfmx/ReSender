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
			<nav class="navbar navbar-expand-lg navbar-light bg-light">
				<a href="/" class="navbar-brand">ReSender</a>
				<div class="collapse navbar-collapse">
					<ul class="navbar-nav mr-auto">
					</ul>
					<form class="form-inline my-2 my-lg-0">
						<input class="form-control mr-sm-2" type="search" placeholder="Search" aria-label="Search">
						<button class="btn btn-outline-success my-2 my-sm-0" type="submit">Search</button>
					</form>	
				</div>
			</nav>
		`	
	}
}


customElements.define("nav-bar", NavBar)
