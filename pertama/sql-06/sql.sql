kite bikin database bernama db_belajar_golang

CREATE TABLE IF NOT EXISTS `tb_student` (
  `id` varchar(5) NOT NULL,
  `name` varchar(255) NOT NULL,
  `age` int(11) NOT NULL,
  `grade` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `tb_student` (`id`, `name`, `age`, `grade`) VALUES
('A1', 'Sam', 18, 1),
('B1', 'Adit', 45, 2),
('C1', 'Dandi', 34, 3),
('D1', 'Rauf', 78, 4),
('E1', 'Gusnur', 5, 4);
('F1', 'Ayat', 123, 4);

ALTER TABLE `tb_student` ADD PRIMARY KEY (`id`);