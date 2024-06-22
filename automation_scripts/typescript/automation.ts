import * as fs from 'fs';

const directoryPath = '/path/to/directory';

fs.readdir(directoryPath, (err, files) => {
  if (err) {
    return console.log('Unable to scan directory: ' + err);
  }
  files.forEach((file) => {
    console.log('File name:', file);
  });
});

console.log('Automation script executed successfully!');
