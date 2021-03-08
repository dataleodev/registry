CREATE TABLE IF NOT EXISTS regions(
    id VARCHAR(50) UNIQUE NOT NULL PRIMARY KEY,
    name VARCHAR (100) UNIQUE,
    description TEXT NOT NULL
);
insert into regions (id, name, description) values ('KGF7032DC41', 'Oak', 'Maecenas pulvinar lobortis est. Phasellus sit amet erat. Nulla tempus. Vivamus in felis eu sapien cursus vestibulum. Proin eu mi. Nulla ac enim. In tempor, turpis nec euismod scelerisque, quam turpis adipiscing lorem, vitae mattis nibh ligula nec sem.');
insert into regions (id, name, description) values ('KYV7796OY49', 'Sunfield', 'Suspendisse accumsan tortor quis turpis. Sed ante. Vivamus tortor. Duis mattis egestas metus. Aenean fermentum. Donec ut mauris eget massa tempor convallis. Nulla neque libero, convallis eget, eleifend luctus, ultricies eu, nibh. Quisque id justo sit amet sapien dignissim vestibulum. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Nulla dapibus dolor vel est.');
insert into regions (id, name, description) values ('SIG3183MQ05', 'Utah', 'Nullam sit amet turpis elementum ligula vehicula consequat. Morbi a ipsum. Integer a nibh. In quis justo. Maecenas rhoncus aliquam lacus. Morbi quis tortor id nulla ultrices aliquet. Maecenas leo odio, condimentum id, luctus nec, molestie sed, justo.');
insert into regions (id, name, description) values ('YUS0591EU89', 'Fieldstone', 'Cras in purus eu magna vulputate luctus.');
insert into regions (id, name, description) values ('KHJ3892WY84', 'Thompson', 'Vestibulum rutrum rutrum neque. Aenean auctor gravida sem. Praesent id massa id nisl venenatis lacinia. Aenean sit amet justo. Morbi ut odio. Cras mi pede, malesuada in, imperdiet et, commodo vulputate, justo. In blandit ultrices enim.');
insert into regions (id, name, description) values ('VTL1860BQ04', 'Darwin', 'Aenean auctor gravida sem.');
insert into regions (id, name, description) values ('XJK7682WX08', 'Harper', 'Morbi non lectus. Aliquam sit amet diam in magna bibendum imperdiet. Nullam orci pede, venenatis non, sodales sed, tincidunt eu, felis. Fusce posuere felis sed lacus. Morbi sem mauris, laoreet ut, rhoncus aliquet, pulvinar sed, nisl. Nunc rhoncus dui vel sem. Sed sagittis. Nam congue, risus semper porta volutpat, quam pede lobortis ligula, sit amet eleifend pede libero quis orci. Nullam molestie nibh in lectus. Pellentesque at nulla.');


CREATE TABLE IF NOT EXISTS nodes(
    uuid VARCHAR(100) UNIQUE NOT NULL PRIMARY KEY,
    addr VARCHAR(100) NOT NULL,
    key VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    type INT NOT NULL,
    region VARCHAR(100) NOT NULL,
    latd VARCHAR(100) NOT NULL,
    long VARCHAR(100) NOT NULL,
    created VARCHAR(100) NOT NULL,
    master VARCHAR(100)
);

CREATE TABLE IF NOT EXISTS users(
    id VARCHAR(100) UNIQUE NOT NULL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    region VARCHAR(100) UNIQUE NOT NULL,
    password VARCHAR(100) UNIQUE NOT NULL,
    created VARCHAR(100) UNIQUE NOT NULL
);