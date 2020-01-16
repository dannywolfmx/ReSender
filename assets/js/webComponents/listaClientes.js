class ListaClientes extends HTMLElement{
	constructor(){
		super()
		let s = this.attachShadow({mode:'open'})
		fetch('clientes').then((r)=>{
			return r.json()
		}).then((clientes)=>{
			s.innerHTML = `
				<ul>
					${clientes.map((cliente)=> `<li>${cliente.nombre}</li>`).join('')}
				</ul>
			`
		}).catch((e)=>{
			console.log("Error al obtener clientes: ",e)
		})
	}
	
}

customElements.define("lista-clientes",ListaClientes)
