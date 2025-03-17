import {
  IoIosCheckmarkCircle,
  IoIosCloseCircle,
  IoIosMailUnread,
} from "react-icons/io";
import CardWrapper from "./card/card-wrapper";
import { Link } from "react-router-dom";
import { Button } from "./ui/button";

interface StatusCardProps {
  type: "success" | "error" | "email";
  headerLabel: string;
  message: string;
  buttonLabel: string;
  buttonHref: string;
}

const StatusCard = ({
  type,
  headerLabel,
  message,
  buttonLabel,
  buttonHref,
}: StatusCardProps) => {
  const statusConfig = {
    success: {
      icon: <IoIosCheckmarkCircle className="w-24 h-24 text-green-500" />,
    },
    error: {
      icon: <IoIosCloseCircle className="w-24 h-24 text-destructive" />,
    },
    email: {
      icon: <IoIosMailUnread className="w-24 h-24 text-green-500" />,
    },
  };

  return (
    <CardWrapper headerLabel={headerLabel}>
      <div className="w-full flex flex-col items-center justify-center space-y-8">
        {statusConfig[type].icon}
        <p className="font-semibold text-neutral-500 text-sm text-center">
          {message}
        </p>
        <Link to={buttonHref} className="w-full">
          <Button variant="submit" className="w-full">
            {buttonLabel}
          </Button>
        </Link>
      </div>
    </CardWrapper>
  );
};

export default StatusCard;
