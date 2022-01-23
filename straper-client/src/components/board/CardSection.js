import React from "react";

const CardSection = ({ Icon, title, children }) => {
  return (
    <div className="p-2">
      <span className="flex">
        <Icon size={25} />
        <span className="px-3 font-semibold">{title}</span>
      </span>
      {children}
    </div>
  );
};

export default CardSection;
