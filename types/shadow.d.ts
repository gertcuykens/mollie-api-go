/**
 * Shadow DOM v1
 *
 * Based on https://www.w3.org/TR/2016/WD-shadow-dom-20160830/
 */

interface DocumentOrShadowRoot {
    getSelection(): Selection | null;
    elementFromPoint(x: number, y: number): Element | null;
    elementsFromPoint(x: number, y: number): Element[];
    readonly activeElement: Element | null;
    readonly styleSheets: StyleSheetList;

    // caretPositionFromPoint(x: number, y: number): CaretPosition | null;
}

interface ShadowRoot extends DocumentOrShadowRoot, DocumentFragment {
  /** Represents the shadow host which hosts this shadow root. */
  readonly host: Element;
  innerHTML: string;
}

interface Document extends DocumentOrShadowRoot{
}

interface Element {
  attachShadow(shadowRootInitDict: ShadowRootInit): ShadowRoot;
  readonly assignedSlot: HTMLSlotElement | null;
  slot: string;
  readonly shadowRoot: ShadowRoot | null;
}

interface ShadowRootInit {
  mode: 'open'|'closed';
  delegatesFocus?: boolean;
}

interface Text {
  readonly assignedSlot: HTMLSlotElement | null;
}

interface HTMLSlotElement extends HTMLElement {
  name: string;
  assignedNodes(options?: AssignedNodesOptions): Node[] ;
}

interface AssignedNodesOptions {
  flatten?: boolean
}

interface EventInit {
  scoped?: boolean;
}

interface Event {
  deepPath(): EventTarget[];
  readonly scoped: boolean;
}
