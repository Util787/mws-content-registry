import React from 'react';

const Input = ({
                   value,
                   onChange,
                   placeholder = '',
                   disabled = false,
                   className = '',
                   ...props
               }) => {
    return (
        <input
            type="text"
            value={value}
            onChange={onChange}
            placeholder={placeholder}
            disabled={disabled}
            className={`w-full px-4 py-3 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500 focus:border-transparent disabled:bg-gray-100 disabled:cursor-not-allowed text-gray-900 placeholder-gray-500 ${className}`}
            {...props}
        />
    );
};

export default Input;