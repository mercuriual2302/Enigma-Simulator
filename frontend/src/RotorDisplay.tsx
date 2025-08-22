import React from "react";

interface Props {
  positions: number[];
  onRotate: (idx: number) => void;
}

export default function RotorDisplay({ positions, onRotate }: Props) {
  return (
    <div className="flex justify-center gap-4 mb-4">
      {positions.map((p, i) => (
        <div
          key={i}
          className="w-16 h-16 rounded-full bg-gray-700 flex items-center justify-center text-xl cursor-pointer select-none"
          onClick={() => onRotate(i)}
        >
          {String.fromCharCode(65 + p)}
        </div>
      ))}
    </div>
  );
}
