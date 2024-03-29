import React, { useState } from "react";
import { AiOutlineClockCircle } from "react-icons/ai";
import CardDialog from "./CardDialog";

const Card = ({ card }) => {
  const [isDialogOpen, setIsDialogOpen] = useState(false);

  const tagColor = () => {
    switch (card.priority) {
      case "LOW":
        return "bg-sky-400";
      case "MEDIUM":
        return "bg-orange-400";
      case "HIGH":
        return "bg-red-500";
    }
  };

  const dateStringToMonthDate = () => {
    const date = new Date(card.due_date);
    const month = date.toLocaleString("default", { month: "short" });
    return month + " " + date.getDate();
  };

  const openCardDialog = async () => {
    setIsDialogOpen(true);
  };

  return (
    <div
      className="flex flex-col bg-white rounded-md p-3"
      onClick={() => openCardDialog()}
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
      <div>
        <div className="bg-indigo-300 flex space-x-1 align-center w-fit rounded-md text-gray-700 py-1 px-3">
          <AiOutlineClockCircle size={20} />
          <span className="text-sm">
            {card.due_date && dateStringToMonthDate()}
          </span>
        </div>
      </div>
      <CardDialog
        open={isDialogOpen}
        closeModal={() => setIsDialogOpen(false)}
        card={card}
      />
    </div>
  );
};

export default Card;
