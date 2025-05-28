import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";
import { Form, FormControl, FormField, FormItem } from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { ArrowUp } from "lucide-react";
import { cn } from "@/lib/utils";

interface ChatInputProps {
  onSend: (msg: string) => void;
  isLoading?: boolean;
  isEmpty?: boolean;
}

const formSchema = z.object({
  message: z.string().min(1, "Message cannot be empty"),
});

type FormValues = z.infer<typeof formSchema>;

const ChatInput = ({ onSend, isLoading, isEmpty }: ChatInputProps) => {
  const form = useForm<FormValues>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      message: "",
    },
  });

  const onSubmit = (data: FormValues) => {
    onSend(data.message);
    form.reset();
  };

  return (
    <div className="w-full flex flex-col justify-center items-center py-10 px-4">
      <h1
        className={cn(
          "text-2xl font-bold mb-8 text-center",
          isEmpty ? "block" : "hidden"
        )}
      >
        What are your plans for today?
      </h1>

      <Form {...form}>
        <form
          onSubmit={form.handleSubmit(onSubmit)}
          className="w-full max-w-2xl flex items-center gap-2 bg-green-50 rounded-full px-6 py-4 shadow-md"
        >
          <FormField
            control={form.control}
            name="message"
            render={({ field }) => (
              <FormItem className="flex-1">
                <FormControl>
                  <Input
                    placeholder="Message Kairos"
                    className="border-none shadow-none focus-visible:ring-0 focus-visible:ring-offset-0 text-sm sm:text-base"
                    disabled={isLoading}
                    {...field}
                  />
                </FormControl>
              </FormItem>
            )}
          />
          <Button
            type="submit"
            size="icon"
            className="rounded-full"
            disabled={isLoading || !form.formState.isValid}
          >
            <ArrowUp size={18} />
          </Button>
        </form>
      </Form>
    </div>
  );
};

export default ChatInput;
