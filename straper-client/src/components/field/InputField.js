import React, { useState } from "react";

const InputField = ({ defaultValue, action }) => {
  const [field, setField] = useState(defaultValue);
  return (
    <div>
      <input
        className="bg-transparent focus:outline-none"
        value={field}
        onChange={(e) => setField(e.currentTarget.value)}
        onBlur={(e) => action(e.currentTarget.value)}
      />
    </div>
  );
};

export default InputField;
