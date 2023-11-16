interface GoodMorning {
  wakeup: string;
  time: number;
}

interface Nighty {
  message: string;
  reps: number;
}

type Morning = Nighty | GoodMorning;

function sendMessage(m: Morning) {
  if ('message' in m) {
    console.log('hello');
  }

  if ('wakeup' in m) {
    console.log('nighty');
  }
}

let goodMorningMessage: GoodMorning = {
  wakeup: 'Good morning!',
  time: 7,
};

let nightyMessage: Nighty = {
  message: 'Good night!',
  reps: 3,
};

sendMessage(goodMorningMessage); // Output: Good morning!
sendMessage(nightyMessage); // Output: Good night!
