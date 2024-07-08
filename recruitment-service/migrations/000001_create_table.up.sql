CREATE TABLE Vacancy (
    id              UUID PRIMARY KEY,         -- Har bir bo'sh ish o'rni uchun yagona identifikator
    title           VARCHAR(255) NOT NULL,      -- Bo'sh ish o'rnining nomi
    description     TEXT,                       -- Bo'sh ish o'rnining batafsil tavsifi
    department_id   VARCHAR(255) NOT NULL,      -- Bo'lim identifikatori (tashqi kalit)
    position_id     VARCHAR(255) NOT NULL,      -- Lavozim identifikatori (tashqi kalit)
    status          VARCHAR(50) NOT NULL,       -- Bo'sh ish o'rnining holati (masalan, ochiq, yopiq)
    created_at      TIMESTAMP DEFAULT NOW(),    -- Bo'sh ish o'rni qachon yaratilganini ko'rsatuvchi vaqt belgisi
    updated_at      TIMESTAMP DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP -- Bo'sh ish o'rni oxirgi marta qachon yangilanganini ko'rsatuvchi vaqt belgisi
    deleted_at      BIGINT DEFAULT 0

);

-- Nom bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_vacancy_title ON Vacancy(title);

-- Bo'lim identifikatori bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_vacancy_department_id ON Vacancy(department_id);

-- Lavozim identifikatori bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_vacancy_position_id ON Vacancy(position_id);

-- Holat bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_vacancy_status ON Vacancy(status);




CREATE TABLE JobApplication (
    id              UUID PRIMARY KEY,         -- Har bir ish uchun ariza uchun yagona identifikator
    candidate_id    VARCHAR(255) NOT NULL,      -- Nomzodning identifikatori (tashqi kalit)
    vacancy_id      VARCHAR(255) NOT NULL,      -- Bo'sh ish o'rnining identifikatori (tashqi kalit)
    resume_id       VARCHAR(255) NOT NULL,      -- Rezyumening identifikatori (tashqi kalit)
    status          VARCHAR(50) NOT NULL,       -- Ish uchun arizaning holati
    created_at      TIMESTAMP DEFAULT NOW(),    -- Ish uchun ariza qachon yaratilganini ko'rsatuvchi vaqt belgisi
    updated_at      TIMESTAMP DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP -- Ish uchun ariza oxirgi marta qachon yangilanganini ko'rsatuvchi vaqt belgisi
    deleted_at      BIGINT DEFAULT 0
       
);

-- Nomzod identifikatori bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_job_application_candidate_id ON JobApplication(candidate_id);

-- Bo'sh ish o'rni identifikatori bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_job_application_vacancy_id ON JobApplication(vacancy_id);

-- Rezyume identifikatori bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_job_application_resume_id ON JobApplication(resume_id);

-- Holat bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_job_application_status ON JobApplication(status);



CREATE TABLE Company (
    id              UUID PRIMARY KEY,         -- Har bir kompaniya uchun yagona identifikator
    name            VARCHAR(255) NOT NULL,      -- Kompaniyaning nomi
    address         TEXT,                       -- Kompaniyaning manzili
    industry        VARCHAR(255),               -- Kompaniyaning faoliyat sohasi
    website         VARCHAR(255),               -- Kompaniyaning veb-sayti
    created_at      TIMESTAMP DEFAULT NOW(),    -- Kompaniya qachon yaratilganini ko'rsatuvchi vaqt belgisi
    updated_at      TIMESTAMP DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP -- Kompaniya oxirgi marta qachon yangilanganini ko'rsatuvchi vaqt belgisi
    deleted_at      BIGINT DEFAULT 0

);

-- Nom bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_company_name ON Company(name);

-- Manzil bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_company_address ON Company(address);

-- Sanoat bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_company_industry ON Company(industry);

-- Veb-sayt bo'yicha tezkor qidirish uchun indeks
CREATE INDEX idx_company_website ON Company(website);
