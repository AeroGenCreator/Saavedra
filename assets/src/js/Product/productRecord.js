document.addEventListener('alpine:init', () => {
  Alpine.data('productRecordComponent', () => ({

  //FIELDS
    name: '',
    desc: '',
    measure: '',
    price: '',
    material: '',
    proveedor: '',
    pMeasureRecords: [],
    materialRecords: [],
    proveedorRecords: [],

  init(many2one, record) {
    console.log(many2one, record)
  },

  allRequired() {
    return ([this.name, this.desc, this.measure, this.price, this.material, this.proveedor].includes(""))
    },

  async goBack() {
    try {
      const res = await SecureFetching("/product", { method: "HEAD" })
      if (!res.ok) {
        throw new Error(res.status)
      }
      window.location.href = "/product"
    } catch (error) {
      throw error
    }
  },

  async goHome() {
    const res = await SecureFetching('/welcome', { method: 'HEAD' })
    if (res.ok) {
      window.location.href = '/welcome'
      return
    }
    await this.logOut()
    alert(res.status === 401 ? 'Sesión expirada' : `Error ${res.status}`)
  },

  async logOut() {
    try {
      const res = await fetch('/login', {
        method: 'PATCH',
        credentials: 'include',
        headers: { 'X-Requested-With': 'jsFrontendComponent' },
      })
      if (res.ok) window.location.href = '/login'
    } catch (error) {
      console.error('No fue posible cerrar sesión:', error)
    }
    },

  }))
})
