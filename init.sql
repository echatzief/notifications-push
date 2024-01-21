CREATE TABLE notifications (
    id SERIAL PRIMARY KEY,
    email TEXT,
    name TEXT,
    message TEXT,
    status TEXT DEFAULT 'not_sent',
    sent_at TIMESTAMP DEFAULT NOW()
);
INSERT INTO notifications(email, name, message) VALUES('test@gmail.com', 'Receiver', 'Welcome to our services');
INSERT INTO notifications(email, name, message) VALUES('test@gmail.com', 'Receiver', 'Welcome to our services');
INSERT INTO notifications(email, name, message) VALUES('test@gmail.com', 'Receiver', 'Welcome to our services');
INSERT INTO notifications(email, name, message) VALUES('test@gmail.com', 'Receiver', 'Welcome to our services');