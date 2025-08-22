import React from "react";
import { Canvas } from "@react-three/fiber";
import { OrbitControls } from "@react-three/drei";

interface Props {
  positions: number[];
  onRotate: (idx: number) => void;
}

function Rotor({ index, pos, onRotate }: { index: number; pos: number; onRotate: (idx:number)=>void }) {
  const angle = -(pos / 26) * Math.PI * 2;
  const colors = ["#e53e3e", "#dd6b20", "#38a169"];
  return (
    <mesh
      position={[index * 1.5 - 1.5, 0, 0]}
      rotation={[0, 0, angle]}
      onClick={() => onRotate(index)}
      castShadow
      receiveShadow
    >
      <cylinderGeometry args={[0.5, 0.5, 1, 32]} />
      <meshStandardMaterial color={colors[index % colors.length]} />
    </mesh>
  );
}

export default function Machine3D({ positions, onRotate }: Props) {
  return (
    <Canvas style={{ height: "200px" }} camera={{ position: [0, 3, 6] }} shadows>
      <ambientLight intensity={0.6} />
      <directionalLight position={[5, 5, 5]} intensity={0.5} />
      {positions.map((p, i) => (
        <Rotor key={i} index={i} pos={p} onRotate={onRotate} />
      ))}
      <OrbitControls />
    </Canvas>
  );
}
