import { cn } from "@/lib/utils";
import { useState } from "react";
import ChatInput from "@/components/chat-input";

interface Message {
  role: "user" | "kairos";
  content: string;
}

const ChatPage = () => {
  const [messages, setMessages] = useState<Message[]>([]);

  const handleSend = (input: string) => {
    setMessages((prev) => [...prev, { role: "user", content: input }]);

    setTimeout(() => {
      setMessages((prev) => [
        ...prev,
        {
          role: "kairos",
          content: `I'm Kairos! You said: "${input}" 😊`,
        },
      ]);
    }, 1000);
  };

  return (
    <div
      className={cn(
        "pl-[280px] min-h-screen flex flex-col items-center",
        messages.length === 0 ? "justify-center" : "justify-start"
      )}
    >
      <div
        className={cn(
          "overflow-y-auto px-4 py-25 max-w-2xl mx-auto space-y-4",
          messages.length === 0 ? "hidden" : "block"
        )}
      >
        {messages.map((msg, index) => (
          <div
            key={index}
            className={cn(msg.role === "user" ? "text-right" : "text-left")}
          >
            <p
              className={cn(
                "inline-block px-4 py-2 rounded-xl text-sm",
                msg.role === "user"
                  ? "bg-green-100 text-green-900"
                  : "bg-gray-100 text-gray-800"
              )}
            >
              {msg.content}
            </p>
          </div>
        ))}
      </div>
      <div
        className={cn(
          "w-full",
          messages.length === 0
            ? "flex flex-col justify-center items-center"
            : "fixed bottom-0"
        )}
      >
        <ChatInput onSend={handleSend} isEmpty={messages.length === 0} />
      </div>
    </div>
  );
};

export default ChatPage;
