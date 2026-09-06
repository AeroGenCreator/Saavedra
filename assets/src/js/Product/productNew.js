document.addEventListener('alpine:init', () => {
  Alpine.data('productNewComponent', () => ({

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

  init(measures, material, proveedor) {
  this.pMeasureRecords = measures
  this.materialRecords = material
  this.proveedorRecords = proveedor
  },

  allRequired() {
    return ([this.name, this.desc, this.measure, this.price, this.material, this.proveedor].includes(""))
    },

  async create() {
    try {
      const objMaterial = this.materialRecords.find(item => item.name.toLowerCase() === this.material.toLowerCase())
      const objProveedor = this.proveedorRecords.find(item => item.name.toLowerCase() === this.proveedor.toLowerCase())
      const values = {
        name: this.name,
        description: this.desc,
        pMeasure: this.measure,
        price: this.price,
        materialId: String(objMaterial.id),
        proveedorId: String(objProveedor.id)
      }
      const res = await SecureFetching("/product/new", {
        method: "POST", body: JSON.stringify(values)
      })
      if (!res.ok) {
        throw new Error(res.status)
      }
      window.location.href = "/product"
    } catch (error) {
      throw error
    }
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
