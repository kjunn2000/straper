import React, { useState } from "react";
import CardDialog from "./CardDialog";

const Card = ({ card }) => {
  const [isDialogOpen, setIsDialogOpen] = useState(false);

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
    <div
      className="flex flex-col bg-white rounded-md p-3"
      onClick={() => setIsDialogOpen(true)}
    >
      {card.priority !== "NO" && (
        <div
          className={
            "text-white rounded-xl px-3 py-1 font-semibold text-sm " +
            tagColor()
          }
        >
          {card.priority}
        </div>
      )}
      <div className="break-all text-sm p-2">{card.title}</div>
      <div className="text-right">People involved</div>
      <CardDialog
        open={isDialogOpen}
        closeModal={() => setIsDialogOpen(false)}
        card={card}
      />
    </div>
  );
};

export default Card;
