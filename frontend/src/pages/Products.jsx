import { useEffect, useState } from "react";
import { getProducts, addProduct } from "../api";

export default function Products() {
  const [products, setProducts] = useState([]);
  const [form, setForm] = useState({ name: "", unit: "", price: "" });

  useEffect(() => {
    load();
  }, []);

  const load = async () => {
    const res = await getProducts();
    setProducts(res.data);
  };

  const submit = async (e) => {
    e.preventDefault();
    await addProduct(form);
    setForm({ name: "", unit: "", price: "" });
    load();
  };

  return (
    <div className="p-6">
      <h1 className="text-2xl font-bold mb-4">Products</h1>

      <form onSubmit={submit} className="flex gap-2 mb-6">
        <input
          placeholder="Name"
          value={form.name}
          onChange={(e) => setForm({ ...form, name: e.target.value })}
          className="border p-2 rounded"
        />
        <input
          placeholder="Unit"
          value={form.unit}
          onChange={(e) => setForm({ ...form, unit: e.target.value })}
          className="border p-2 rounded"
        />
        <input
          placeholder="Price"
          type="number"
          step="0.01"
          value={form.price}
          onChange={(e) => setForm({ ...form, price: e.target.value })}
          className="border p-2 rounded"
        />
        <button className="bg-green-600 text-white px-4 py-2 rounded">
          Add
        </button>
      </form>

      <table className="w-full border-collapse border border-gray-400">
        <thead>
          <tr className="bg-gray-200">
            <th className="border px-4 py-2">ID</th>
            <th className="border px-4 py-2">Name</th>
            <th className="border px-4 py-2">Unit</th>
            <th className="border px-4 py-2">Price</th>
          </tr>
        </thead>
        <tbody>
          {products.map((p) => (
            <tr key={p.ID}>
              <td className="border px-4 py-2">{p.ID}</td>
              <td className="border px-4 py-2">{p.Name}</td>
              <td className="border px-4 py-2">{p.Unit}</td>
              <td className="border px-4 py-2">${p.Price}</td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}