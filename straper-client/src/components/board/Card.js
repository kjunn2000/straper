import React from "react";

const Card = ({ card }) => {
  console.log(card);

  const tagColor = () => {
    switch (card.priority) {
      case "LOW":
        return "bg-sky-400";
      case "MEDIUM":
        return "bg-orange-500";
      case "HIGH":
        return "bg-red-600";
    }
  };

  return (
    <div className="flex flex-col bg-white rounded-md p-3 m-3">
      {card.priority !== "NO" && (
        <div
          className={
            "text-white rounded-xl p-1 font-semibold text-sm " + tagColor()
          }
        >
          {card.priority}
        </div>
      )}
      <div className="break-all text-sm p-2">{card.title}</div>
      <div className="text-right">People involved</div>
    </div>
  );
};

export default Card;
