// Define the quiz questions and answers
const questions = [
 {
    question:"In essence, what does 'Due Process of Law' mean",
    answers:[
        {text:"The principle of natural justice", correct:false},
        {text:"The procedure established by law", correct:false},
        {text:"Fair application of law", correct:true},
        {text:"Equality before law", correct:false},
    ]

 },
 {
    question: "Which of the following words were added in the Preamble by the 42nd Constitutional Amendment Act? 1. Socialist 2. Integrity 3. Fraternity Select the correct answer using the code given below?",
    answers:[
        {text:"1 and 3 only", correct:false},
        {text:"1 and 2 only", correct:true},
        {text:"2 and 3 only", correct:false},
        {text:"1, 2 and 3 ", correct:false},
    ]
 },
 {
    question:"A constitutional government by definition is:",
    answers:[
        {text:"A nominated government", correct:false},
        {text:"A multi-party Government", correct:false},
        {text:"A limited government", correct:true},
        {text:"None of the Above", correct:false},
    ]
 },
 {
    question:"The idea of ‘Justice’ in the preamble of the Constitution of India has been borrowed from:",
    answers:[
        {text:"Russian Revolution", correct:true},
        {text:"French Revoltuion", correct:false},
        {text:"American Revolution", correct:false},
        {text:"Japanese Revoltuion", correct:false},
    ]
 },
 {
    question:"Which among the following articles of Constitution of India abolishes the untouchablity?",
    answers:[
        {text:"Article 15", correct:false},
        {text:"Article 16", correct:false},
        {text:"Article 17", correct:true},
        {text:"Article 18", correct:false},
    ]
 }
];

// Define DOM elements and variables
const questionElement = document.getElementById("question");
const answerButtons= document.getElementById("answer-buttons");
const nextButton = document.getElementById("next-btn");

let currentQuestionIndex = 0;// Index to track the current question
let score = 0;  // Track the score

// Function to start the quiz
function startQuiz(){
    currentQuestionIndex = 0; // Reset the question index
    score = 0; // Reset the score
    nextButton.innerHTML = "Next"; // Set text for Next button
    showQuestion();  // Call function to display the first question
}

// Function to display the current question
function showQuestion(){
    resetState();
    let currentQuestion = questions[currentQuestionIndex];
    let questionNo = currentQuestionIndex + 1;
    questionElement.innerHTML = questionNo + "." + currentQuestion.question;  
    
    // Loop through all possible answers and create buttons for them
    currentQuestion.answers.forEach(answer => {
        const button = document.createElement("button");
        button.innerHTML = answer.text;
        button.classList.add("btn");
        answerButtons.appendChild(button);
        if(answer.correct){
            button.dataset.correct = answer.correct;
        }
        button.addEventListener("click", selectAnswer);
    });
}

// Function to reset the answer button states
function resetState(){
    nextButton.style.display = "none";
    while(answerButtons.firstChild){
        answerButtons.removeChild(answerButtons.firstChild);
    }
}


// Function that gets triggered when an answer is selected
function selectAnswer(e){
    const selectedBtn = e.target;
    const isCorrect = selectedBtn.dataset.correct === "true";
    if(isCorrect){
        selectedBtn.classList.add("correct");
        score++;
    }else{
        selectedBtn.classList.add("incorrect");
    }
    // Disable all answer buttons and highlight the correct answer
    Array.from(answerButtons.children).forEach(button => {
        if(button.dataset.correct === "true"){
            button.classList.add("correct");
        }
        button.disabled = true;
    });
    nextButton.style.display = "block";

}
// Function to display the user's score at the end of the quiz
function showScore(){
    resetState();
    questionElement.innerHTML = `Your Score ${score} out of ${questions.length}!`;
    nextButton.innerHTML = "Play Again";
    nextButton.style.display ="block";
}
// Function to handle the "Next" button logic
function handleNextButton(){
    currentQuestionIndex++;
    if(currentQuestionIndex < questions.length){
        showQuestion();
    }else{
        showScore();
    }

}

// Event listener for the "Next" button click
nextButton.addEventListener("click", () => {
    if(currentQuestionIndex < questions.length){
        handleNextButton();
    }else{
        startQuiz();
    }

});

startQuiz();