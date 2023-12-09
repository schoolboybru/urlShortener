import axios from "axios";
import { useState } from "react";
import { useMutation } from "react-query";

export const UrlForm = () => {
  const [shortenedUrl, setShortenedUrl] = useState<string>();

  const baseUrl: string = import.meta.env.VITE_BASE_URL;

  const submitUrl = async ({ value }: { value: string }) => {
    const response = await axios.post(`${baseUrl}/url?value=${value}`);
    return response.data;
  };

  const { mutate, isLoading, error } = useMutation(submitUrl);

  const handleSubmit = (event: React.SyntheticEvent) => {
    event.preventDefault();

    const target = event.target as typeof event.target & {
      value: { value: string };
    };
    const url = target.value;

    mutate(url);
  };

  return (
    <div className="flex bg-slate-900 h-screen flex-col items-center justify-center">
      <form
        className="bg-white shadow-md rounded px-8 pt-6 pb-8 mb-4"
        onSubmit={handleSubmit}
      >
        <div>
          <h1 className="text-3xl font-bold ">Enter a Url to Shorten</h1>
        </div>
        <div className="flex items-center border-b border-teal-500 py-2">
          <input
            className="appearance-none bg-transparent border-none w-full text-gray-700 mr-3 py-1 leading-tight focus:outline-none"
            type="text"
            placeholder="Url"
            name="value"
          ></input>
          <button
            className="flex-shrink-0 bg-teal-500 hover:bg-teal-700 border-teal-500 hover:border-teal-700 text-sm border-4 text-white py-1 px-2 rounded"
            type="submit"
          >
            Submit
          </button>
          {!isLoading && !error && shortenedUrl && <text></text>}
        </div>
      </form>
    </div>
  );
};
