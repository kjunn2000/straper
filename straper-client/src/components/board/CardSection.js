import React from "react";

const CardSection = ({ Icon, title }) => {
  return (
    <div>
      <span className="flex">
        <Icon />
        {title}
      </span>
    </div>
  );
};

export default CardSection;
