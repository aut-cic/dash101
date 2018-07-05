/*
 * +===============================================
 * | Author:        Parham Alvani (parham.alvani@gmail.com)
 * |
 * | Creation Date: 03-08-2017
 * |
 * | File Name:     toc.js
 * +===============================================
 */
const toc = {
  initToCOnLoad: function (titles) {
    let sections = document.getElementsByClassName('toc')

    for (let section of sections) {
      const titlesListElement = document.createElement('ul');

      for (let i = 0; i < titles.length; i++) {
        let title = titles[i]
        let node = document.createElement('li');

        if (section.dataset.selected && (section.dataset.selected === title || section.dataset.selected == i)) {
          node.className += 'material-select'
        }

        node.appendChild(document.createTextNode(title));
        titlesListElement.appendChild(node);
      }

      section.appendChild(titlesListElement);
    }
  }
}
