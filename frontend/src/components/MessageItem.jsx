import { useEffect, useRef } from 'react'

export default function MessageItem({ type, message, title, timeout = 3000, onRemove }) {

  const ref = useRef();

  useEffect(() => {
    const timer = setTimeout(() => {
      ref.current?.addEventListener("animationend", onRemove, { once: true });
      ref.current?.classList.add("message-hide");
    }, timeout);

    return () => clearTimeout(timer);
  }, [onRemove]);

  return (
  <div ref={ref} className={`message message-${type}`}>
    {title} - 
    {message}
  </div>
  );
}
