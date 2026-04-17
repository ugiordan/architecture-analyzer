// Diagram toolbar: zoom in/out, fullscreen, and reset controls
// Appears on hover over any mermaid diagram, similar to GitHub's image viewer
(function () {
  var ZOOM_STEP = 0.25;
  var MIN_SCALE = 0.5;
  var MAX_SCALE = 3;

  function createToolbar(container) {
    if (container.querySelector(".diagram-toolbar")) return;

    var svg = container.querySelector("svg");
    if (!svg) return;

    var toolbar = document.createElement("div");
    toolbar.className = "diagram-toolbar";

    var scale = 1;
    var isFullscreen = false;

    function makeBtn(label, title, onClick) {
      var btn = document.createElement("button");
      btn.textContent = label;
      btn.title = title;
      btn.addEventListener("click", function (e) {
        e.stopPropagation();
        e.preventDefault();
        onClick();
      });
      return btn;
    }

    function applyScale() {
      svg.style.transform = "scale(" + scale + ")";
      svg.style.transformOrigin = "top left";
    }

    var zoomIn = makeBtn("+", "Zoom in", function () {
      scale = Math.min(scale + ZOOM_STEP, MAX_SCALE);
      applyScale();
    });

    var zoomOut = makeBtn("\u2212", "Zoom out", function () {
      scale = Math.max(scale - ZOOM_STEP, MIN_SCALE);
      applyScale();
    });

    var reset = makeBtn("1:1", "Reset zoom", function () {
      scale = 1;
      svg.style.transform = "";
      svg.style.transformOrigin = "";
    });

    var fullscreenBtn = makeBtn("\u26F6", "Fullscreen", function () {
      isFullscreen = !isFullscreen;
      container.classList.toggle("fullscreen", isFullscreen);
      fullscreenBtn.textContent = isFullscreen ? "\u2715" : "\u26F6";
      fullscreenBtn.title = isFullscreen ? "Exit fullscreen" : "Fullscreen";
      if (!isFullscreen) {
        scale = 1;
        svg.style.transform = "";
        svg.style.transformOrigin = "";
      }
    });

    toolbar.appendChild(zoomIn);
    toolbar.appendChild(zoomOut);
    toolbar.appendChild(reset);
    toolbar.appendChild(fullscreenBtn);

    container.style.position = "relative";
    container.appendChild(toolbar);
  }

  function initToolbars() {
    document.querySelectorAll("pre.mermaid, .mermaid").forEach(function (el) {
      createToolbar(el);
    });
  }

  // Run after mermaid renders (needs delay for mkdocs-material)
  if (document.readyState === "loading") {
    document.addEventListener("DOMContentLoaded", function () {
      setTimeout(initToolbars, 1500);
    });
  } else {
    setTimeout(initToolbars, 1500);
  }

  // Re-init on page navigation (mkdocs-material instant loading)
  document.addEventListener("DOMContentSwitch", function () {
    setTimeout(initToolbars, 1500);
  });

  // Watch for dynamically rendered diagrams
  var observer = new MutationObserver(function (mutations) {
    var hasMermaid = false;
    for (var i = 0; i < mutations.length; i++) {
      var added = mutations[i].addedNodes;
      for (var j = 0; j < added.length; j++) {
        if (added[j].nodeType === 1 && (added[j].matches && (added[j].matches("pre.mermaid svg, .mermaid svg") || added[j].querySelector && added[j].querySelector("pre.mermaid svg, .mermaid svg")))) {
          hasMermaid = true;
        }
      }
    }
    if (hasMermaid) initToolbars();
  });
  observer.observe(document.body, { childList: true, subtree: true });

  // ESC exits fullscreen
  document.addEventListener("keydown", function (e) {
    if (e.key === "Escape") {
      document.querySelectorAll(".fullscreen").forEach(function (el) {
        el.classList.remove("fullscreen");
        var btn = el.querySelector(".diagram-toolbar button:last-child");
        if (btn) {
          btn.textContent = "\u26F6";
          btn.title = "Fullscreen";
        }
        var svg = el.querySelector("svg");
        if (svg) {
          svg.style.transform = "";
          svg.style.transformOrigin = "";
        }
      });
    }
  });
})();
