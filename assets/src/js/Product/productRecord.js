document.addEventListener('alpine:init', () => {
  Alpine.data('productRecordComponent', (m2o, record) => ({

  //FIELDS
    id: record.id || '',
    name: record.name || '',
    desc: record.description || '',
    measure: record.pMeasure || '',
    price: record.price || '',
    material: record.material || '',
    proveedor: record.proveedor || '',
    pMeasureRecords: m2o.pMeasureRecords || [],
    materialRecords: m2o.materialRecords || [],
    proveedorRecords: m2o.proveedorRecords || [],

  allRequired() {
    return ([this.name, this.desc, this.measure, this.price, this.material, this.proveedor].includes(""))
    },

    async updateRecord() {
      try {
        const objMaterial = this.materialRecords.find(item => item.name.toLowerCase() === this.material.toLowerCase())
        const objProveedor = this.proveedorRecords.find(item => item.name.toLowerCase() === this.proveedor.toLowerCase())
        const values = {
          id: String(this.id),
          name: this.name,
          description: this.desc,
          pMeasure: this.measure,
          price: this.price,
          materialId: String(objMaterial.id),
          proveedorId: String(objProveedor.id)
        }
        const res = await SecureFetching('/product/record', {
          method: 'PUT', body: JSON.stringify(values),
        })
        if (!res.ok) throw new Error(await res.text())
        window.location.href = "/product"
      } catch (error) {
        throw error
      }
    },

    async deleteRecord() {
      try {
        const res = await SecureFetching(`/product/record?id=${this.id}`, {
          method: 'DELETE', body: JSON.stringify({ id: this.id })
        })
        if (!res.ok) throw new Error(await res.text())
        window.location.href = '/product'
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
