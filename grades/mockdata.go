package grades

func init() {
	students = []Student{
		{
			ID:        1,
			FirstName: "Averill",
			LastName:  "Simon",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 85,
				},
				{
					Title: "Week 1 Homework",
					Type:  GradeHowmework,
					Score: 94,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		{
			ID:        2,
			FirstName: "Marge",
			LastName:  "Garrard",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 100,
				},
				{
					Title: "Week 1 Homework",
					Type:  GradeHowmework,
					Score: 100,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 88,
				},
			},
		},
		{
			ID:        3,
			FirstName: "Sydney",
			LastName:  "Barber",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 77,
				},
				{
					Title: "Week 1 Homework",
					Type:  GradeHowmework,
					Score: 0,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 65,
				},
			},
		},
		{
			ID:        4,
			FirstName: "Loui",
			LastName:  "Easton",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 88,
				},
				{
					Title: "Week 1 Homework",
					Type:  GradeHowmework,
					Score: 93,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 84,
				},
			},
		},
		{
			ID:        5,
			FirstName: "Kylee",
			LastName:  "Attwood",
			Grades: []Grade{
				{
					Title: "Quiz 1",
					Type:  GradeQuiz,
					Score: 95,
				},
				{
					Title: "Week 1 Homework",
					Type:  GradeHowmework,
					Score: 99,
				},
				{
					Title: "Quiz 2",
					Type:  GradeQuiz,
					Score: 100,
				},
			},
		},
	}
}
