import { useState } from "react";
import { Encrypt, Configure } from "../wailsjs/go/main/App";

const rotorOptions = ["I", "II", "III", "IV", "V"];

export default function App() {
  const [rotors, setRotors] = useState<string[]>(["I", "II", "III"]);
  const [positions, setPositions] = useState<number[]>([0, 0, 0]);
  const [reflector, setReflector] = useState("B");
  const [plugs, setPlugs] = useState("");
  const [plain, setPlain] = useState("");
  const [cipher, setCipher] = useState("");

  const updateRotor = (idx: number, value: string) => {
    const next = [...rotors];
    next[idx] = value;
    setRotors(next);
  };

  const updatePosition = (idx: number, value: string) => {
    const next = [...positions];
    next[idx] = Math.max(0, Math.min(25, parseInt(value) || 0));
    setPositions(next);
  };

  const handleEncrypt = async () => {
    const plugPairs = plugs
      .split(/\s+/)
      .filter((p) => p.length === 2)
      .map((p) => p.toUpperCase());
    await Configure(rotors, positions, reflector, plugPairs);
    const result = await Encrypt(plain.toUpperCase());
    setCipher(result);
  };

  return (
    <div className="h-full bg-gradient-to-br from-gray-900 via-purple-900 to-indigo-900 text-white flex flex-col">
      <header className="text-center py-4 text-3xl font-bold drop-shadow-lg">
        Enigma Simulator
      </header>
      <main className="flex-grow flex items-center justify-center p-4">
        <div className="bg-white/10 backdrop-blur-md rounded-lg p-6 w-full max-w-4xl shadow-2xl">
          <div className="grid grid-cols-3 gap-4 mb-6">
            {rotors.map((r, i) => (
              <div key={i} className="flex flex-col">
                <label className="mb-1 text-sm font-semibold">
                  Rotor {i + 1}
                </label>
                <select
                  className="text-black rounded px-2 py-1 mb-2"
                  value={r}
                  onChange={(e) => updateRotor(i, e.target.value)}
                >
                  {rotorOptions.map((opt) => (
                    <option key={opt}>{opt}</option>
                  ))}
                </select>
                <input
                  type="number"
                  min="0"
                  max="25"
                  className="text-black rounded px-2 py-1"
                  value={positions[i]}
                  onChange={(e) => updatePosition(i, e.target.value)}
                />
              </div>
            ))}
          </div>
          <div className="flex flex-col md:flex-row gap-4 mb-6">
            <div className="flex-1">
              <label className="block mb-1 text-sm font-semibold">Reflector</label>
              <select
                className="text-black rounded px-2 py-1 w-full"
                value={reflector}
                onChange={(e) => setReflector(e.target.value)}
              >
                <option>B</option>
                <option>C</option>
              </select>
            </div>
            <div className="flex-1">
              <label className="block mb-1 text-sm font-semibold">
                Plugboard (pairs e.g. AB CD)
              </label>
              <input
                className="text-black rounded px-2 py-1 w-full"
                value={plugs}
                onChange={(e) => setPlugs(e.target.value)}
              />
            </div>
          </div>
          <div className="mb-4">
            <label className="block mb-1 text-sm font-semibold">Plaintext</label>
            <textarea
              className="text-black rounded w-full p-2 h-24"
              value={plain}
              onChange={(e) => setPlain(e.target.value)}
            />
          </div>
          <div className="mb-4">
            <label className="block mb-1 text-sm font-semibold">Ciphertext</label>
            <textarea
              className="text-black rounded w-full p-2 h-24 bg-gray-200"
              value={cipher}
              readOnly
            />
          </div>
          <button
            className="w-full py-2 bg-gradient-to-r from-pink-500 to-yellow-400 rounded font-semibold text-black hover:from-pink-400 hover:to-yellow-300 transition-colors"
            onClick={handleEncrypt}
          >
            Encrypt
          </button>
        </div>
      </main>
    </div>
  );
}
