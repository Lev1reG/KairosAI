import React from "react";

interface InputFieldProps {
    icon: React.ReactNode;
    placeholder: string;
    type?: string;
}

const InputField: React.FC<InputFieldProps> = ({ icon, placeholder, type = "text" }) => {
    return (
        <div className="flex items-center bg-gray-100 rounded-full px-4 py-2">
            <span className="mr-3">{icon}</span>
            <input
                type={type}
                placeholder={placeholder}
                className="bg-transparent flex-1 outline-none"
            />
        </div>
    );
};

export default InputField;