class ListaClientes extends HTMLElement{
	constructor(){
		super()
		this.attachShadow({mode:'open'})
		fetch('clientes').then((r)=>{
			return r.json()
		}).then(this.template.bind(this)).catch((e)=>{
			console.log("Error al obtener clientes: ",e)
		})
	}
	
	//Imprime la lista de clientes
	template(clientes){
		 this.shadowRoot.innerHTML= `
			<ul>
				${clientes.map((cliente)=> `<li>${cliente.nombre}</li>`).join('')}
			</ul>
		`
	}
	
}

customElements.define("lista-clientes",ListaClientes)
